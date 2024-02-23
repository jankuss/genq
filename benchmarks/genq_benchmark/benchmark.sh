hyperfine --parameter-step-size 1 --parameter-scan count 1 10 --prepare './reset.sh && dart ./bin/gen.dart {count}' '../../tool/genq --input ./lib' --export-json ./benchmark.json
