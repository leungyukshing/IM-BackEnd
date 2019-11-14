#!/usr/bin/env bash

BIN_ROOT=$(cd $(dirname ${BASH_SOURCE[0]}); pwd)
PROJECT_ROOT=$(cd ${BIN_ROOT}/..; pwd)

rm ${PROJECT_ROOT}/coverage*.out 2> /dev/null
cd ${PROJECT_ROOT}/docker
EXIT_CODE=0


if [[ $? -eq 0 ]]
then
  DOCKER_COMPOSE="env docker-compose -f docker-compose.ci.yml -p ci"
  ${DOCKER_COMPOSE} up -d mysql

  until ${DOCKER_COMPOSE} exec -T mysql mysql -uroot -pci -e "show processlist"; do
    sleep 10;
  done

  ${DOCKER_COMPOSE} run --rm ci bash ./bin/run-ci
  EXIT_CODE=$?
fi

if [[ ${EXIT_CODE} == 0 ]];then
  echo "SUCCESS"
else
  echo "FAIL"
fi