package main

import (
	"context"
	"dagger/foo/internal/dagger"
	"fmt"
)

type Foo struct{}

func (f *Foo) TestCacheVolumePersistence(ctx context.Context, input string) (string, error) {
	_, err := f.PopulateCache(ctx, input)
	if err != nil {
		return "", err
	}

	output, err := f.ListCache(ctx)
	if err != nil {
		return "", err
	}

	return output, nil
}

func (f *Foo) PopulateCache(ctx context.Context, input string) (*dagger.Container, error) {
	return dag.Container().From("alpine:latest").
		WithMountedCache("/foo-cache", dag.CacheVolume("foo-cache")).
		WithExec([]string{"sh", "-c", fmt.Sprintf("echo '%s' > /foo-cache/bar.txt", input)}).
		Sync(ctx)
}

func (f *Foo) ListCache(ctx context.Context) (string, error) {
	return dag.
		Container().
		From("alpine:latest").
		WithMountedCache("/foo-cache", dag.CacheVolume("foo-cache")).
		WithExec([]string{"sh", "-c", "cat /foo-cache/bar.txt"}).
		Stdout(ctx)
}

func (f *Foo) WithNewFile(ctx context.Context) (*dagger.Container, error) {
	return dag.
		Container().
		From("alpine:latest").
		WithMountedCache("/foo-cache", dag.CacheVolume("foo-cache")).
		WithNewFile("/foo-cache/somefile.txt", "contents of somefile").
		Sync(ctx)
}
