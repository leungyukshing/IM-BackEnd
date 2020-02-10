#!/usr/bin/env bash

BIN_ROOT=$(cd $(dirname ${BASH_SOURCE[0]}); pwd)
PROJECT_ROOT=$(cd ${BIN_ROOT}/..; pwd)

rm ${PROJECT_ROOT}/coverage*.out 2> /dev/null
cd ${PROJECT_ROOT}/docker
EXIT_CODE=0


if [[ $? -eq 0 ]]
then
  DOCKER_COMPOSE="env docker-compose -f docker-compose.yml -p ci"
  ${DOCKER_COMPOSE} up -d mysql ci

  until ${DOCKER_COMPOSE} exec -T mysql mysql -uroot -pci -e "show processlist"; do
    sleep 10;
  done

  ${DOCKER_COMPOSE} run --rm ci bash ./bin/run-ci.sh
  EXIT_CODE=$?
  # ${DOCKER_COMPOSE} down -v
fi

if [[ ${EXIT_CODE} == 0 ]];then
  echo "SUCCESS"
  exit 0
else
  echo "FAIL"
  exit 1
fi