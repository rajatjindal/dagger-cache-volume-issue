package main

import (
	"context"
	"dagger/foo/internal/dagger"
	"fmt"
	"strings"
)

type Foo struct{}

func (f *Foo) TestCacheVolumePersistence(ctx context.Context) error {
	_, err := f.PopulateCache(ctx)
	if err != nil {
		return err
	}

	output, err := f.ListCache(ctx)
	if err != nil {
		return err
	}

	if !strings.Contains(output, "bar.txt") {
		return fmt.Errorf("%#v does not contain bar.txt file", output)
	}

	return nil
}

func (f *Foo) PopulateCache(ctx context.Context) (*dagger.Container, error) {
	return dag.Container().From("alpine:latest").
		WithMountedCache("/foo-cache", dag.CacheVolume("foo-cache")).
		WithExec([]string{"sh", "-c", "echo 'i am in cache 123456' > /foo-cache/bar.txt"}).Terminal().
		Sync(ctx)
}

func (f *Foo) ListCache(ctx context.Context) (string, error) {
	return dag.
		Container().
		From("alpine:latest").
		WithMountedCache("/foo-cache", dag.CacheVolume("foo-cache")).
		WithExec([]string{"sh", "-c", "cat /foo-cache/bar.txt"}).Terminal().
		Stdout(ctx)
}
