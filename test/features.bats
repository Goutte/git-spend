#!/usr/bin/env bats

# https://github.com/bats-core/bats-core
# Run:
#     bats test

setup() {
    load 'test_helper/bats-support/load'
    load 'test_helper/bats-assert/load'

    TESTS_DIR="$( cd "$( dirname "$BATS_TEST_FILENAME" )" >/dev/null 2>&1 && pwd )"
    PROJECT_DIR="$( dirname "$TESTS_DIR" )"
    gitime=${PROJECT_DIR}/build/gitime

    # We copy this project into a temporary dir (in RAM),
    # check out the appropriate fixture-XX tag,
    # and run integration testing on that tmp repo.
    TMP_FIXTURE_DIR="/tmp/gitime-test"
    rm -rf $TMP_FIXTURE_DIR

    cp -R "$PROJECT_DIR" "$TMP_FIXTURE_DIR"
    cd "$TMP_FIXTURE_DIR" || exit

    git switch -c fixture-00
}

@test "gitime" {
  run $gitime
  [ "$status" -eq 0 ]
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

@test "gitime sum --author notfound" {
  run $gitime sum --author notfound
  # shouldn't we fail, here?   TBD
  #assert_failure
  assert_success  # …meanwhile
  assert_output "No time-tracking directives /spend or /spent found in commits."
}

