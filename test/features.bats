#!/usr/bin/env bats

# https://github.com/bats-core/bats-core
# Run:
#     make test-acceptance

# We use gitime's own repo as fixture.
# We copy this project into a temporary fixture dir (in RAM),
# and then have it check out the appropriate fixture-XX tag,
# and finally run integration testing on that temporary repo.
TMP_FIXTURE_DIR="/tmp/gitime-test"

setup() {
    load 'test_helper/bats-support/load'
    load 'test_helper/bats-assert/load'

    TESTS_DIR="$( cd "$( dirname "$BATS_TEST_FILENAME" )" >/dev/null 2>&1 && pwd )"
    PROJECT_DIR="$( dirname "$TESTS_DIR" )"
    COVERAGE_DIR=${PROJECT_DIR}/test-coverage
    gitime=${PROJECT_DIR}/build/gitime

    if [ "$GITIME_COVERAGE" == "1" ] ; then
      echo "Setting up coverage in ${COVERAGE_DIR}"
      mkdir -p "${COVERAGE_DIR}"
      export GOCOVERDIR=${COVERAGE_DIR}
      gitime="${gitime}-coverage"
    fi

    export GITIME_NO_STDIN=1
    export TZ="Europe/Paris"

    cp -R "$PROJECT_DIR" "$TMP_FIXTURE_DIR"
    cd "$TMP_FIXTURE_DIR" || exit

    git stash
    git checkout tags/fixture-00 -b fixture-00
    echo "success: ignore the unable to rmdir warning above (benign)"

    git log fixture-00 > fixture-00.log
    git log 0.1.0 > 0.1.0.log
}

teardown() {
    rm -rf $TMP_FIXTURE_DIR
    rm -f fixture-00.log
    rm -f 0.1.0.log
}

@test "gitime" {
  run $gitime
  assert_success
  assert_output --partial 'Gather information about /spent time from commit messages'
}

@test "gitime hohohoooo should fail" {
  run $gitime hohohoooo
  assert_failure
}

@test "gitime help sum" {
  run $gitime help sum
  assert_success
}

@test "gitime sum --help" {
  run $gitime sum --help
  assert_success
}

@test "gitime sum" {
  run $gitime sum
  assert_success
  assert_output "1 week 3 hours"
}

@test "gitime sum --minutes" {
  run $gitime sum --minutes
  assert_success
  assert_output "2580"
}

@test "gitime sum --hours" {
  run $gitime sum --hours
  assert_success
  assert_output "43"
}

@test "gitime sum --days" {
  run $gitime sum --days
  assert_success
  assert_output "5"
}

@test "gitime sum --weeks" {
  run $gitime sum --weeks
  assert_success
  assert_output "1"
}

@test "gitime sum --months" {
  run $gitime sum --months
  assert_success
  assert_output "0"
}

@test "gitime sum --author Goutte" {
  run $gitime sum --author Goutte
  assert_success
  assert_output "1 week 3 hours"
}

@test "gitime sum --author antoine@goutenoir.com" {
  run $gitime sum --author antoine@goutenoir.com
  assert_success
  assert_output "1 week 3 hours"
}

@test "gitime sum --author notfound (should fail)" {
  run $gitime sum --author notfound
  # shouldn't we fail, here?   TBD
  #assert_failure
  assert_success  # …meanwhile
  assert_output "No time-tracking directives /spend or /spent found in commits."
}

@test "gitime sum --since <commit>" {
  run $gitime sum --since 786a30642fe37368b0b65cbca8ca1a5c4b6c97b8
  assert_success
  assert_output "1 day 6 hours 3 minutes"
}

@test "gitime sum --since <commit short>" {
  run $gitime sum --since 786a3064
  assert_success
  assert_output "1 day 6 hours 3 minutes"
}

@test "gitime sum --since <tag>" {
  run $gitime sum --since 0.2.0
  assert_success
  assert_output "1 day 6 hours 3 minutes"
}

@test "gitime sum --since <absent commit short>" {
  run $gitime sum --since caca999
  assert_failure
}

@test "gitime sum --since <wrong> (should fail)" {
  run $gitime sum --since lololololo
  assert_failure
}

@test "gitime sum --until <commit>" {
  run $gitime sum --until 786a30642fe37368b0b65cbca8ca1a5c4b6c97b8
  assert_success
  assert_output "3 days 4 hours 57 minutes"
}

@test "gitime sum --until <commit short>" {
  run $gitime sum --until 786a3064
  assert_success
  assert_output "3 days 4 hours 57 minutes"
}

@test "gitime sum --until <tag>" {
  run $gitime sum --until 0.2.0
  assert_success
  assert_output "3 days 4 hours 57 minutes"
}

@test "gitime sum --until <absent commit short>" {
  run $gitime sum --until caca666
  assert_failure
}

@test "gitime sum --until <wrong>" {
  run $gitime sum --until trololololo
  assert_failure
}

@test "gitime sum --until 0.1.0" {
  run $gitime sum --until 0.1.0
  assert_success
  assert_output "1 day 7 hours 57 minutes"
}

@test "gitime sum --until tags/<tag>" {
  run $gitime sum --until 0.1.0
  assert_success
  assert_output "1 day 7 hours 57 minutes"
}

@test "gitime sum --since 0.1.0" {
  run $gitime sum --since 0.1.0
  assert_success
  assert_output "3 days 3 hours 3 minutes"
}

@test "gitime sum --since 0.1.0 --until 0.1.1" {
  run $gitime sum --since 0.1.0 --until 0.1.1
  assert_success
  assert_output "30 minutes"
}

@test "gitime sum --since <date>" {
  # Sun Mar 26 22:11:03 2023 of 4527140510c2b77a9f2a6eb947b5391d4e2173a9
  run $gitime sum --since 2023-03-27
  assert_success
  assert_output "2 hours"
}

@test "gitime sum --since <date time>" {
  run $gitime sum --since "2023-03-26 22:15:00"
  assert_success
  assert_output "2 hours 1 minute"
  # We'd want, but no cigar ; time parsing in Golang is quite weird
  #run $gitime sum --since "2023-03-26 22:15"
  #assert_success
}

@test "gitime sum --since <date rfc3339>" {
  run $gitime sum --since 2023-03-26T22:15:00Z
  assert_success
  assert_output "2 hours 1 minute"
}

@test "gitime sum --until <date>" {
  run $gitime sum --until 2023-03-25
  assert_success
  assert_output "1 day 3 hours 55 minutes"
}

@test "gitime sum --since <date> --until <date>" {
  run $gitime sum --since "2023-03-25 03:30:00" --until "2023-03-25 13:37:00"
  assert_success
  assert_output "2 hours 15 minutes"
}

@test "gitime sum does not accept mixed dates and refs in ranges" {
  run $gitime sum --until 2023-03-27 --since 0.1.0
  assert_failure
  run $gitime sum --since 2023-03-24 --until 0.2.0
  assert_failure
}

@test "gitime sum using stdin" {
  export GITIME_NO_STDIN=0
  run bash -c "cat fixture-00.log | $gitime sum"
#  run $gitime sum < fixture-00.log
  assert_success
  assert_output "1 week 3 hours"
}

@test "gitime sum using another stdin" {
  export GITIME_NO_STDIN=0
  run bash -c "cat 0.1.0.log | $gitime sum"
#  run bash -c "$gitime sum < 0.1.0.log"
  assert_success
  assert_output "1 day 7 hours 57 minutes"
}

@test "gitime sum using stdin does not accept --no-merges" {
  export GITIME_NO_STDIN=0
  run bash -c "cat fixture-00.log | $gitime sum --no-merges"
#  run bash -c "$gitime sum --no-merges < fixture-00.log"
  assert_failure
}

@test "gitime sum using stdin does not accept --author" {
  export GITIME_NO_STDIN=0
  run bash -c "cat fixture-00.log | $gitime sum --author Goutte"
#  run bash -c "$gitime sum --author Goutte < fixture-00.log"
  assert_failure
}

@test "gitime sum using stdin does not accept --since" {
  export GITIME_NO_STDIN=0
  run bash -c "cat fixture-00.log | $gitime sum --since 0.1.0"
#  run bash -c "$gitime sum --since 0.1.0 < fixture-00.log"
  assert_failure
}

@test "gitime sum using stdin does not accept --until" {
  export GITIME_NO_STDIN=0
  run bash -c "cat fixture-00.log | $gitime sum --until 0.1.0"
#  r0un bash -c "$gitime sum --until 0.1.0 < fixture-00.log"
  assert_failure
}