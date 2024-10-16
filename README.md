repro.sh

```
#!/bin/bash

export BARHOME=/Users/rajatjindal/go/src/github.com/rajatjindal/dagger-cache-volume-issue/bar
export FOOHOME=/Users/rajatjindal/go/src/github.com/rajatjindal/dagger-cache-volume-issue
export DAGGERHOME=/Users/rajatjindal/go/src/github.com/dagger/dagger
export WITHDEV=$DAGGERHOME/hack/with-dev

## this won't work
## withMountedCache(cache: {cache(key: \"xxh3:3f44273d9c319a53volume-name\"): CacheVolume!}
cd $BARHOME
$WITHDEV dagger call use-cache-volume-name --name volume-name

## this will populate cache
## withMountedCache(cache: {cache(key: \"xxh3:45843aaf90e01497volume-name\"): CacheVolume!}
cd $FOOHOME
ID=`SHUTUP=1 $WITHDEV dagger call -q get-cache-volume-id`

## still wont work
## withMountedCache(cache: {cache(key: \"xxh3:3f44273d9c319a53volume-name\"): CacheVolume!}
cd $BARHOME
$WITHDEV dagger call use-cache-volume-name --name volume-name

## this will call bar function from within foo and it works
## cache: {cache(key: \"xxh3:45843aaf90e01497volume-name\"): CacheVolume!}
cd $FOOHOME
$WITHDEV dagger call use-cache-volume-across-module-using-id --id $ID


## now calling from inside bar works. it should not have worked.
cd $BARHOME
$WITHDEV dagger call use-cache-volume-name --name volume-name
```
