export MQTT_BROKER=tcp://127.0.0.1:1883
export HUB_URL=http://127.0.0.1:5151

echo "Setting up services"
docker-compose --project-name ioh_tester up --build -d
echo "Waiting for things to start"
sleep 10

echo "Running tests"
cd tester
go test -v || FAILED=1
cd ..

echo "Cleaning up"
docker-compose --project-name ioh_tester down

if [[ -n $FAILED ]]
then
  echo FAILED
  exit 1
fi
echo OK
