# zkFil-lib

[![Build Status](https://travis-ci.org/xuxinlai2002/zkFilzkFil-lib.svg?branch=master)](https://travis-ci.org/xuxinlai2002/zkFilzkFil-lib)

zkFil-lib is the underlying core library for [zkFil system](https://github.com/xuxinlai2002/zkFilzkFil-node). It fully implements PoD (proof of delivery) protocol and also provides a CLI interface together with Golang bindings.

zkFil-lib has four main parts as followed.

- [`zk_setup`](zk_setup/) - Generates public parameters of system
- [`zk_publish`](zk_publish/) - Publish utility to preprocess data and calculate authenticators
- [`zk_core`](zk_core/) - Core implementations of different zkFil trading features
- [`zk_go`](zk_go/) - Golang bindings for `zk_setup` and `zk_core` functions

## Related zkFil projects

- [zkFil-node](https://github.com/xuxinlai2002/zkFilzkFil-node) Node application written in Golang for sellers (Alice) and buyers (Bob). It deals with communication, smart contract calling, data transferring, and other zkFil protocol interactions.
- [zkFil-contract](https://github.com/xuxinlai2002/zkFilzkFil-contract) Smart contracts for zkFil _Decentralized Exchange_.

## Dependencies

**We would recommend Ubuntu 18.04 for convenience. The default GCC on Ubuntu 16.04 is outdated and you may encounter some annoying issues.**

- GCC 7.3 or above (***If you upgrade GCC manually from an old version like GCC5, you should re-install the boost library.***)
- Go 1.12
- Boost 1.69.0 or newer
- OpenMP 5.3.1 or newer

```shell
# For Ubuntu
sudo apt-get install libomp-dev
sudo apt-get install libcrypto++-dev
sudo apt-get install libboost-all-dev
sudo apt-get install libgmp-dev

# For macOS
brew install libcryptopp
brew install boost
brew install gmp
```

## Build

```shell
# Download zkFil-lib code
mkdir zkFil && cd zkFil
git clone https://github.com/xuxinlai2002/zkFilzkFil-lib.git

# Pull libsnark submodule
cd zkFil-lib
git submodule init && git submodule update
cd depends/libsnark
git submodule init && git submodule update

# Build libsnark
mkdir build && cd build
# - On Ubuntu
cmake -DCMAKE_INSTALL_PREFIX=../../install -DMULTICORE=ON -DWITH_PROCPS=OFF -DWITH_SUPERCOP=OFF -DCURVE=MCL_BN128 ..
# - Or on macOS ref https://github.com/scipr-lab/libsnark/issues/99#issuecomment-367677834
CPPFLAGS=-I/usr/local/opt/openssl/include LDFLAGS=-L/usr/local/opt/openssl/lib PKG_CONFIG_PATH=/usr/local/opt/openssl/lib/pkgconfig cmake -DCMAKE_INSTALL_PREFIX=../../install -DMULTICORE=OFF -DWITH_PROCPS=OFF -DWITH_SUPERCOP=OFF -DCURVE=MCL_BN128 -DUSE_ASM=OFF ..
make && make install

# Build zkFil-lib
cd ../../..
make

# These files should be generated after successful build.
# zk_setup/zk_setup
# zk_publish/zk_publish
# zk_core/libzk_core.so
# zk_core/zk_core
```

## Usage

### zk_setup

```shell
$ ./zk_setup -h

command line options:
  -h [ --help ]                         Use -h or --help to list all arguments
  -o [ --output_path ] arg (=zksnark_key)
                                        Provide the output path
  -c [ --count ] arg (=1024)            Provide the count
  -v [ --verbose ] arg (=0)             Enable libff log
```

Use `zk_setup` to generate zkFil zkSNARK public parameters.

```shell
./zk_setup
```

`zksnark_key` is generated after a successful setup.

### zk_publish

```shell
$ ./zk_publish -h
command line options:
  -h [ --help ]                         Use -h or --help to list all arguments
  ---d data_dir -m table -f file -o output_dir -t table_type -k keys
                                        publish table file
  ---d data_dir -m plain -f file -o output_dir -c column_num
                                        publish plain file
  -d [ --data_dir ] arg (=.)            Provide the configure file dir
  -m [ --mode ] arg (=plain)            Provide pod mode (plain, table)
  -f [ --publish_file ] arg             Provide the file which want to publish
  -o [ --output_dir ] arg               Provide the publish path
  -t [ --table_type ] arg (=csv)        Provide the publish file type in table
                                        mode (csv)
  -c [ --column_num ] arg (=1023)       Provide the column number per
                                        block(line) in plain mode (default
                                        1023)
  -k [ --vrf_colnum_index ] arg         Provide the publish file vrf key column
                                        index positions in table mode (for
                                        example: -k 0 1 3)
  -u [ --unique_key ] arg               Provide the flag if publish must unique
                                        the key in table mode (for example: -u
                                        1 0 1)
  --omp_thread_num arg (=0)             Provide the number of the openmp
                                        thread, 1: disable openmp, 0: default.
```

Use `zk_publish` to preprocess target data.

zkFil supports two modes: binary mode and table mode.

- binary mode
- table mode (CSV files)

```shell
# binary mode
./zk_publish -m plain -f test.txt -o plain_data -c 1024

# table mode
./zk_publish -m table -f test1000.csv -o table_data -t csv -k 0 1
```

Check the output folder after publishing.

### zk_core

```shell
$ ./zk_core -h
command line options:
  -h [ --help ]              Use -h or --help to list all arguments
  -d [ --data_dir ] arg (=.) Provide the configure file dir
  -m [ --mode ] arg          Provide pod mode (plain, table)
  -a [ --action ] arg        Provide action (range_pod, ot_range_pod,
                             vrf_query, ot_vrf_query...)
  -p [ --publish_dir ] arg   Provide the publish dir
  -o [ --output_dir ] arg    Provide the output dir
  --demand_ranges arg        Provide the demand ranges
  --phantom_ranges arg       Provide the phantom range(plain mode)
  -k [ --query_key ] arg     Provide the query key name(table mode)
  -v [ --key_value ] arg     Provide the query key values(table mode, for
                             example -v value_a value_b value_c)
  -n [ --phantom_key ] arg   Provide the query key phantoms(table mode, for
                             example -n phantoms_a phantoms_b phantoms_c)
  --omp_thread_num arg       Provide the number of the openmp thread, 1:
                             disable openmp, 0: default.
  -c [ --use_c_api ]
  --test_evil
  --dump_ecc_pub
```

`zk_core` supports several mode combinations. We have `atomic-swap`, `atomic-swap-vc` and `complaint` trade mode for binary and table files. Moreover, we could employ _oblivious transfer_, `OT` mode, for privacy-preserving download. Furthermore, we are allowed to do `vrf_query` of structured table data, which could be combined with `OT` mode for the private query.

```shell
./zk_core -m plain -a complaint_pod -p plain_data -o plain_output --demand_ranges 0-2

./zk_core -m plain -a atomic_swap_pod -p plain_data -o plain_output --demand_ranges 1-10

./zk_core -m plain -a atomic_swap_pod_vc -p plain_data -o plain_output --demand_ranges 1-10

./zk_core -m table -a ot_vrf_query -p table_data -o table_output -k "Emp ID" -v 313736 964888 abc -n 350922 aaa eee bbb
```

Check [here](zk_core/README.md) for more CLI interface examples. You could look over [each](zk_core/scheme_atomic_swap_test.cc) [test](zk_core/scheme_atomic_swap_vc_test.cc) for detailed protocol implementation.

### zk_go

A simple Golang wrapper for zkFil-lib is provided for easier library integration.

```shell
cd zk_go
export GO111MODULE=on
make test
```

All tests should pass as expected. You could check over [tests](zk_go/plain/api_test.go) to learn about usage.

## Learn more?

+ White paper: an overview introduction of the zkFil system.
+ [Technical paper](https://xuxinlai2002/zkFil.github.io/zkFil-node/paper.pdf): a document with theoretic details to those who are interested in the theory we are developing.
+ Community: join us on [*Discord*](https://discord.gg/tfUH886) and follow us on [*Twitter*](https://twitter.com/SECBIT_IO) please!

## License

zkFil-lib is released under the terms of the MIT license. See LICENSE for more information or see https://opensource.org/licenses/MIT.
