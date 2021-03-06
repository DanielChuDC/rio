// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type EdsSnapshot struct {
	Upstreams UpstreamList
}

func (s EdsSnapshot) Clone() EdsSnapshot {
	return EdsSnapshot{
		Upstreams: s.Upstreams.Clone(),
	}
}

func (s EdsSnapshot) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	if _, err := s.hashUpstreams(hasher); err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (s EdsSnapshot) hashUpstreams(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Upstreams.AsInterfaces()...)
}

func (s EdsSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	hasher := fnv.New64()
	UpstreamsHash, err := s.hashUpstreams(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("upstreams", UpstreamsHash))
	snapshotHash, err := s.Hash(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return append(fields, zap.Uint64("snapshotHash", snapshotHash))
}

type EdsSnapshotStringer struct {
	Version   uint64
	Upstreams []string
}

func (ss EdsSnapshotStringer) String() string {
	s := fmt.Sprintf("EdsSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s EdsSnapshot) Stringer() EdsSnapshotStringer {
	snapshotHash, err := s.Hash(nil)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return EdsSnapshotStringer{
		Version:   snapshotHash,
		Upstreams: s.Upstreams.NamespacesDotNames(),
	}
}
