#!/usr/bin/env bats

# Acceptance test suite, made with BATS.
# https://github.com/bats-core/bats-core
# Run:
#     make test-acceptance

# We use git-spend's own repo as fixture for tests. (🐕 woof)
# We copy this project into a temporary fixture dir (in RAM),
# and then have it check out the appropriate fixture-XX tag,
# and finally run integration testing on that temporary repo.
# See the setup() BATS hook defined at the bottom of this file.
# THIS DIRECTORY WILL BE `RM -RF` SO BEWARE OF WHAT'S IN THERE.
TMP_FIXTURE_DIR="/tmp/git-spend-fixture"

@test "git-spend" {
  run $git_spend
  assert_success
  assert_output --partial 'Manage time-tracking /spent directives in commit messages'
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
  assert_output --partial 'The /spend and /spent directives will be parsed and summed'
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
  assert_output --partial "No time-tracking /spend directives found in commits"
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
  run $git_spend sum --until tags/0.1.0
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
  # This can still be done, and would be nice, but I find my solution to be … inelegant. #mr-welcome
  #run $git_spend sum --since "2023-03-26 22:15"
  #run $git_spend sum --since "2023-03"  # and perhaps this as well?
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

@test "git-spend sum but nothing was found" {
  run $git_spend sum --since "2023-03-25" --until "2023-03-25"
  assert_success
  assert_output --partial "No time-tracking /spend directives found in commits"
}

@test "git-spend sum does not accept mixed dates and refs in ranges" {
  run $git_spend sum --until 2023-03-27 --since 0.1.0
  assert_failure
  run $git_spend sum --since 2023-03-24 --until 0.2.0
  assert_failure
}

@test "git-spend sum ignores stdin by default" {
  run bash -c "cat 0.1.0.log | $git_spend sum"
  assert_success
  assert_output "1 week 3 hours"  # and not "1 day 7 hours 57 minutes" for 0.1.0
}

@test "git-spend sum --stdin using |" {
  run bash -c "cat fixture-00.log | $git_spend sum --stdin"
  assert_success
  assert_output "1 week 3 hours"

  run bash -c "cat 0.1.0.log | $git_spend sum --stdin"
  assert_success
  assert_output "1 day 7 hours 57 minutes"  # and not "1 week 3 hours"
}

@test "git-spend sum --stdin using <" {
  run bash -c "$git_spend sum --stdin < fixture-00.log"
  assert_success
  assert_output "1 week 3 hours"

  run bash -c "$git_spend sum --stdin < 0.1.0.log"
  assert_success
  assert_output "1 day 7 hours 57 minutes"
}

@test "git-spend sum --stdin does not accept --target" {
  run bash -c "cat fixture-00.log | $git_spend sum --stdin --target ${PROJECT_DIR}"
  assert_failure
}

@test "git-spend sum --stdin does not accept --no-merges" {
  run bash -c "cat fixture-00.log | $git_spend sum --stdin --no-merges"
  assert_failure
}

@test "git-spend sum --stdin does not accept --author" {
  run bash -c "cat fixture-00.log | $git_spend sum --stdin --author Goutte"
  assert_failure
}

@test "git-spend sum --stdin does not accept --since" {
  run bash -c "cat fixture-00.log | $git_spend sum --stdin --since 0.1.0"
  assert_failure
}

@test "git-spend sum --stdin does not accept --until" {
  run bash -c "cat fixture-00.log | $git_spend sum --stdin --until 0.1.0"
  assert_failure
}

@test "LANGUAGE=fr_FR git-spend" {
  export LANGUAGE=fr_FR
  run $git_spend
  assert_success
  assert_output --partial 'Gérer les directives /spend inscrites dans les messages de commit'
}

@test "LANGUAGE=fr git-spend" {
  export LANGUAGE=fr
  run $git_spend
  assert_success
  assert_output --partial 'Gérer les directives /spend inscrites dans les messages de commit'
}

@test "LC_ALL=fr_FR git-spend" {
  export LC_ALL=fr_FR
  run $git_spend
  assert_success
  assert_output --partial 'Gérer les directives /spend inscrites dans les messages de commit'
}

@test "LANG=fr_FR git-spend" {
  export LANG=fr_FR
  run $git_spend
  assert_success
  assert_output --partial 'Gérer les directives /spend inscrites dans les messages de commit'
}

@test "LANG=fr git-spend sum" {
  export LANG=fr
  run $git_spend sum --until tags/0.1.0
  assert_success
  assert_output '1 jour 7 heures 57 minutes'
}

@test "LC_ALL has priority over LANG" {
  export LANG=en_US
  export LC_ALL=fr_FR

  run $git_spend
  assert_success
  assert_output --partial 'Gérer les directives /spend inscrites dans les messages de commit'
}

@test "LANGUAGE has priority over LC_ALL" {
  export LANGUAGE=fr_FR
  export LC_ALL=en_US

  run $git_spend
  assert_success
  assert_output --partial 'Gérer les directives /spend inscrites dans les messages de commit'
}

@test "Default language is english" {
  export LANGUAGE=""
  export LC_ALL=""
  export LANG=""

  run $git_spend
  assert_success
  assert_output --partial 'Manage time-tracking /spent directives in commit messages'
}

@test "Unhandled language falls back to english" {
  export LANGUAGE="es_CL"
  run $git_spend
  assert_success
  assert_output --partial 'Manage time-tracking /spent directives in commit messages'
}

@test "Unhandled locale should fallback to default locale if language is handled" {
  export LANGUAGE="fr_CA"
  run $git_spend
  assert_success
  assert_output --partial 'Gérer les directives /spend inscrites dans les messages de commit'
}

@test "Generate man pages" {
  run $git_spend man
  assert_success
}

@test "Installing man pages requires sudo" {
  skip  # hehe, CI is root, I forgot ; coverage will lower, but that's OK

  run $git_spend man --output /usr/local/share/man/man8
  assert_failure
  assert_output --partial 'permission denied'

  run $git_spend man --output /usr/share/man/man8
  assert_failure
  assert_output --partial 'permission denied'

  run $git_spend man --install
  assert_failure
  assert_output --partial 'permission denied'
}

# ---

setup() {
    load 'test_helper/bats-support/load'
    load 'test_helper/bats-assert/load'
    export TZ="Europe/Paris"

    TESTS_DIR="$( cd "$( dirname "$BATS_TEST_FILENAME" )" >/dev/null 2>&1 && pwd )"
    PROJECT_DIR="$( dirname "$TESTS_DIR" )"
    COVERAGE_DIR="${PROJECT_DIR}/test-coverage"
    git_spend="${PROJECT_DIR}/build/git-spend"

    cd "${PROJECT_DIR}" || exit

    if [ "$GIT_SPEND_COVERAGE" == "1" ] ; then
      echo "Setting up coverage in ${COVERAGE_DIR}"
      mkdir -p "${COVERAGE_DIR}"
      export GOCOVERDIR=${COVERAGE_DIR}
      git_spend="${git_spend}-coverage"
    fi

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
