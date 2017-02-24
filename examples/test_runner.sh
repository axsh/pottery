#!/bin/bash

work_dir="{{.GitRepositoryURI}}"
testcase_dir="${work_dir}/{{.TestCaseDirectoryName}}"
result_file="${work_dir}/test_result.txt"
log_file="${work_dir}/test_log.txt"

\rm -f ${result_file} ${log_file}

cd ${testcase_dir}

find . -type f -name "*_server.sh" | while read server_script; do
  prefix="`echo ${server_script} | awk -F'_server' '{print $1}'`"
  echo "----- ${prefix} -----" >> ${log_file}
  client_script="${prefix}_client.sh"
  chmod 755 ${server_script} ${client_script}
  server_log="`${server_script}`"
  server_script_result=$?
  echo "+++++ server_log +++++" >> ${log_file}
  echo "${server_log}" >> ${log_file}
  if [ ${server_script_result} -ne 0 ]; then
    echo "${prefix},NG" >> ${result_file}
  else
    sleep 2
    client_log="`${client_script}`"
    client_script_result=$?
    echo "+++++ client_log +++++" >> ${log_file}
    echo "${client_log}" >> ${log_file}
    if [ ${client_script_result} -ne 0 ]; then
      echo "${prefix},NG" >> ${result_file}
    else
      echo "${prefix},OK" >> ${result_file}
    fi
  fi
  echo "---------------------" >> ${log_file}
done
