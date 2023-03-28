#!/usr/bin/env bats

# Acceptance test suite, made with BATS.
# https://github.com/bats-core/bats-core
# Run:
#     make test-acceptance

# We use git-spend's own repo as fixture for tests.
# We copy this project into a temporary fixture dir (in RAM),
# and then have it check out the appropriate fixture-XX tag,
# and finally run integration testing on that temporary repo.
# See the setup() BATS hook defined at the bottom of this file.
# THIS DIRECTORY WILL BE `RM -RF` SO BEWARE OF WHAT'S IN HERE.
TMP_FIXTURE_DIR="/tmp/git-spend-test"

@test "git-spend" {
  run $git_spend
  assert_success
  assert_output --partial 'Gather information about /spent time from commit messages'
}

@test "git-spend hohohoooo should fail" {
  run $git_spend hohohoooo
  assert_failure
}

@test "git-spend help sum" {
  run $git_spend help sum
  assert_success
}

@test "git-spend sum --help" {
  run $git_spend sum --help
  assert_success
}

@test "git-spend sum" {
  run $git_spend sum
  assert_success
  assert_output "1 week 3 hours"
}

@test "git-spend sum --target <dir>" {
  cd "${PROJECT_DIR}"
  run $git_spend sum --target "${TMP_FIXTURE_DIR}"
  assert_success
  assert_output "1 week 3 hours"
}

@test "git-spend sum --target <404 dir> should fail" {
  run $git_spend sum --target "/to/code/or/not/to/code"
  assert_failure
}

@test "git-spend sum --minutes" {
  run $git_spend sum --minutes
  assert_success
  assert_output "2580"
}

@test "git-spend sum --hours" {
  run $git_spend sum --hours
  assert_success
  assert_output "43"
}

@test "git-spend sum --days" {
  run $git_spend sum --days
  assert_success
  assert_output "5"
}

@test "git-spend sum --weeks" {
  run $git_spend sum --weeks
  assert_success
  assert_output "1"
}

@test "git-spend sum --months" {
  run $git_spend sum --months
  assert_success
  assert_output "0"
}

@test "git-spend sum unit formats are mutually exclusive" {
  run $git_spend sum --months --days
  assert_failure
  run $git_spend sum --hours --minutes --weeks
  assert_failure
}

@test "git-spend sum --author Goutte" {
  run $git_spend sum --author Goutte
  assert_success
  assert_output "1 week 3 hours"
}

@test "git-spend sum --author antoine@goutenoir.com" {
  run $git_spend sum --author antoine@goutenoir.com
  assert_success
  assert_output "1 week 3 hours"
}

@test "git-spend sum --author notfound (should fail)" {
  run $git_spend sum --author notfound
  # shouldn't we fail, here?   TBD
  #assert_failure
  assert_success  # …meanwhile
  assert_output "No time-tracking directives /spend or /spent found in commits."
}

@test "git-spend sum --since <commit>" {
  run $git_spend sum --since 786a30642fe37368b0b65cbca8ca1a5c4b6c97b8
  assert_success
  assert_output "1 day 6 hours 3 minutes"
}

@test "git-spend sum --since <commit short>" {
  run $git_spend sum --since 786a3064
  assert_success
  assert_output "1 day 6 hours 3 minutes"
}

@test "git-spend sum --since <tag>" {
  run $git_spend sum --since 0.2.0
  assert_success
  assert_output "1 day 6 hours 3 minutes"
}

@test "git-spend sum --since <absent commit short>" {
  run $git_spend sum --since caca999
  assert_failure
}

@test "git-spend sum --since <wrong> (should fail)" {
  run $git_spend sum --since lololololo
  assert_failure
}

@test "git-spend sum --until <commit>" {
  run $git_spend sum --until 786a30642fe37368b0b65cbca8ca1a5c4b6c97b8
  assert_success
  assert_output "3 days 4 hours 57 minutes"
}

@test "git-spend sum --until <commit short>" {
  run $git_spend sum --until 786a3064
  assert_success
  assert_output "3 days 4 hours 57 minutes"
}

@test "git-spend sum --until <tag>" {
  run $git_spend sum --until 0.2.0
  assert_success
  assert_output "3 days 4 hours 57 minutes"
}

@test "git-spend sum --until <absent commit short>" {
  run $git_spend sum --until caca666
  assert_failure
}

@test "git-spend sum --until <wrong>" {
  run $git_spend sum --until trololololo
  assert_failure
}

@test "git-spend sum --until 0.1.0" {
  run $git_spend sum --until 0.1.0
  assert_success
  assert_output "1 day 7 hours 57 minutes"
}

@test "git-spend sum --until tags/<tag>" {
  run $git_spend sum --until 0.1.0
  assert_success
  assert_output "1 day 7 hours 57 minutes"
}

@test "git-spend sum --since 0.1.0" {
  run $git_spend sum --since 0.1.0
  assert_success
  assert_output "3 days 3 hours 3 minutes"
}

@test "git-spend sum --since 0.1.0 --until 0.1.1" {
  run $git_spend sum --since 0.1.0 --until 0.1.1
  assert_success
  assert_output "30 minutes"
}

@test "git-spend sum --since <date>" {
  run $git_spend sum --since 2023-03-27
  assert_success
  assert_output "2 hours"
}

@test "git-spend sum --since <date time>" {
  run $git_spend sum --since "2023-03-26 22:15:00"
  assert_success
  assert_output "2 hours 1 minute"
  # Want to tolerate missing minutes, but no cigar ; time parsing in Golang is quite peculiar
  #run $git_spend sum --since "2023-03-26 22:15"
  #assert_success
}

@test "git-spend sum --since <date rfc3339>" {
  run $git_spend sum --since 2023-03-26T22:15:00Z
  assert_success
  assert_output "2 hours 1 minute"
}

@test "git-spend sum --until <date>" {
  run $git_spend sum --until 2023-03-25
  assert_success
  assert_output "1 day 3 hours 55 minutes"
}

@test "git-spend sum --since <date> --until <date>" {
  run $git_spend sum --since "2023-03-25 03:30:00" --until "2023-03-25 13:37:00"
  assert_success
  assert_output "2 hours 15 minutes"
}

@test "git-spend sum does not accept mixed dates and refs in ranges" {
  run $git_spend sum --until 2023-03-27 --since 0.1.0
  assert_failure
  run $git_spend sum --since 2023-03-24 --until 0.2.0
  assert_failure
}

@test "git-spend sum without stdin" {
  #export GIT_SPEND_NO_STDIN=0
  run bash -c "cat /dev/null | $git_spend sum"
  assert_success
  assert_output "1 week 3 hours"
}

@test "git-spend sum using stdin" {
  export GIT_SPEND_NO_STDIN=0
  run bash -c "cat fixture-00.log | $git_spend sum"
#  run $git_spend sum < fixture-00.log
  assert_success
  assert_output "1 week 3 hours"
}

@test "git-spend sum using another stdin" {
  export GIT_SPEND_NO_STDIN=0
  run bash -c "cat 0.1.0.log | $git_spend sum"
#  run bash -c "$git_spend sum < 0.1.0.log"
  assert_success
  assert_output "1 day 7 hours 57 minutes"
}

@test "git-spend sum using stdin does not accept --no-merges" {
  export GIT_SPEND_NO_STDIN=0
  run bash -c "cat fixture-00.log | $git_spend sum --no-merges"
#  run bash -c "$git_spend sum --no-merges < fixture-00.log"
  assert_failure
}

@test "git-spend sum using stdin does not accept --author" {
  export GIT_SPEND_NO_STDIN=0
  run bash -c "cat fixture-00.log | $git_spend sum --author Goutte"
#  run bash -c "$git_spend sum --author Goutte < fixture-00.log"
  assert_failure
}

@test "git-spend sum using stdin does not accept --since" {
  export GIT_SPEND_NO_STDIN=0
  run bash -c "cat fixture-00.log | $git_spend sum --since 0.1.0"
#  run bash -c "$git_spend sum --since 0.1.0 < fixture-00.log"
  assert_failure
}

@test "git-spend sum using stdin does not accept --until" {
  export GIT_SPEND_NO_STDIN=0
  run bash -c "cat fixture-00.log | $git_spend sum --until 0.1.0"
#  r0un bash -c "$git_spend sum --until 0.1.0 < fixture-00.log"
  assert_failure
}

# ---

setup() {
    load 'test_helper/bats-support/load'
    load 'test_helper/bats-assert/load'

    TESTS_DIR="$( cd "$( dirname "$BATS_TEST_FILENAME" )" >/dev/null 2>&1 && pwd )"
    PROJECT_DIR="$( dirname "$TESTS_DIR" )"
    COVERAGE_DIR=${PROJECT_DIR}/test-coverage
    git_spend=${PROJECT_DIR}/build/git-spend

    cd "${PROJECT_DIR}" || exit

    if [ "$git_spend_COVERAGE" == "1" ] ; then
      echo "Setting up coverage in ${COVERAGE_DIR}"
      mkdir -p "${COVERAGE_DIR}"
      export GOCOVERDIR=${COVERAGE_DIR}
      git_spend="${git_spend}-coverage"
    fi

    # CI buffers unwanted data in stdin, so let's just disable stdin for most tests
    export GIT_SPEND_NO_STDIN=1
    export TZ="Europe/Paris"

    cp -R "${PROJECT_DIR}" "${TMP_FIXTURE_DIR}"
    cd "${TMP_FIXTURE_DIR}" || exit

    git stash
    git checkout tags/fixture-00 -b fixture-00
    echo "success: ignore the unable to rmdir warning above (benign)"

    git log > fixture-00.log
    git log 0.1.0 > 0.1.0.log
}

teardown() {
    rm -rf $TMP_FIXTURE_DIR
    rm -f fixture-00.log
    rm -f 0.1.0.log
}
