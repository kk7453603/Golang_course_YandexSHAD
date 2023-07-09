[Русский](syllabus.md)
# Syllabus

1. Introduction. Course syllabus. Course artifacts, assessment criteria.
   Philosophy of design. if, switch, for. Hello, world. Command
   line arguments. Word count. Animated gif. Fetching URL. Fetching
   URL concurrently. Web server. Tour of go. Local IDE
   setup. Submitting solutions to automated
   grading. gofmt. goimports. linting. Submitting PR's with bug fixes.
   
   Exercises:
   - [sum](../sum/README-en.md)
   - [tour0](../tour0/README-en.md)
   - [wordcount](../wordcount/README.md)
   - [urlfetch](../urlfetch/README.md)
   - [fetchall](../fetchall/README.md)

2. Language basics. names, declarations, variables,
   assignments. type declarations. packages and files. scope. Zero
   value. Memory allocation. Stack vs heap. Basic data
   types. Constants. Composite data types. Arrays. Slices. Maps. Structs.
   JSON. text/template. string and []byte. Working with unicode. Unicode replacements character.
   Functions. Variadic functions. Anonymous functions. Errors.

   Exercises:
   - [hotelbusiness](../hotelbusiness/README.md)
   - [hogwarts](../hogwarts/README.md)
   - [utf8](../utf8/README.md)
   - [varfmt](../varfmt/README.md)
   - [speller](../speller/README.md)
   - [forth](../forth/README.md)
3. Methods. Value receiver vs pointer receiver. Embedding. Method
   value. Encapsulation. Interfaces. Interface as a contract. io.Writer, io.Reader with
   implementations. sort.Interface. error. http.Handler. Interface as enum. Type assertion. Type switch. The bigger the
   interface, the weaker the abstraction. Error handling. panic,
   defer, recover. errors.{Unwrap,Is,As}. fmt.Errorf. %w.
   
   Exercises:
   - [otp](../otp/README.md)
   - [lrucache](../lrucache/README.md)
   - [externalsort](../externalsort/README.md)
   - [retryupdate](../retryupdate/README.md)
   - [ciletters](../ciletters/README.md)
4. Goroutines and channels. clock server. echo server. Channel size. Blocking/non-blocking reading. select
   statement. Channel axioms. `time.After`. `time.NewTicker`. Pipeline
   pattern. Cancellation. Parallel loop. sync.WaitGroup. Error handling in concurrent code. errgroup.Group. Concurrent web
   crawler. Concurrent directory traversal.
   Exercises:
   - [tour1](../tour1/README.md)
   - [once](../once/README.md)
   - [rwmutex](../rwmutex/README.md)
   - [waitgroup](../waitgroup/README.md)
   - [cond](../cond/README.md)
   - [ratelimit](../ratelimit/README.md)
5. Advanced testing. Subtests. *testing.B. (*T).Logf. (*T).Skipf. (*T).FailNow.
   testing.Short(), testing flags. Mock generation. testify/{require,assert}. testify/suite. Test fixture.
   Integration tests. Goroutine leak detector. TestingMain. Coverage. Benchmarks comparison.
   
   Exercises:
   - [testequal](../testequal/README.md)
   - [fileleak](../fileleak/README.md)
   - [tabletest](../tabletest/README.md)
   - [tparallel](../tparallel/README.md)
   - [iprange](../iprange/README.md)
6. Concurrency with shared memory. sync.Mutex. sync.RWMutex. sync.Cond. atomic. sync.Once.
   Race detector. Async cache. Package context. context.WithTimeout.

   Exercises:
   - [dupcall](../dupcall/README.md)
   - [keylock](../keylock/README.md)
   - [batcher](../batcher/README.md)
   - [pubsub](../pubsub/README.md)
7. net/http. Passing request-scoped data. http middleware. chi.Router. Graceful server shutdown.

   Exercises:
   - [urlshortener](../urlshortener/README.md)
   - [digitalclock](../digitalclock/README.md)
   - [middleware](../middleware/README.md)
   - [olympics](../olympics/README.md)
   - [firewall](../firewall/README.md)
8. database/sql, sqlx, working with databases, redis.
   
   Exercises:
   - [dao](../dao/README.md)
   - [ledger](../ledger/README.md)
   - [shopfront](../shopfront/README.md)
   - [wscat](../wscat/README.md)

9. Generics. Generics features. Type params. Constraints. Type-checking. Type sets. Comparable interface.
   Type interface. Inference. 
   
   Exercises:
   - [genericsum](../genericsum/README.md)
   - [treeiter](../treeiter/README.md)
   - [coverme](../coverme/README.md)

10. Reflection. reflect.Type and reflect.Value. struct tags. net/rpc. encoding/gob.
    sync.Map. reflect.DeepEqual.
   
    Exercises:
    - [reversemap](../reversemap/README.md)
    - [jsonlist](../jsonlist/README.md)
    - [jsonrpc](../jsonrpc/README.md)
    - [structtags](../structtags/README.md)

11. IO package, Reader and Writer implementations from standard lib.
    Low-level programming. unsafe. Package binary. bytes.Buffer. cgo, syscall.
    
    Exercises:
    - [illegal](../illegal/README.md)
    - [blowfish](../blowfish/README.md)

12. Garbage collector (GC) design. Write barrier. Stack growth. GC pause. GOGC. sync.Pool. Goroutine scheduler. GOMACPROCS. Threads leak. 
    Go tooling. pprof. CPU and Memory profiling. Cross-compiling. GOOS, GOARCH. CGO_ENABLED=0.
    Build tags. go modules. godoc. Code generation.

    Exercises:
    - [gzep](../gzep/README.md)

13. go/types and x/analysis. Go code static analysis.

    Exercises:
    - [testifycheck](../testifycheck/README.md)

Bonus exercises:
- [consistenthash](../consistenthash/README.md)
- [gossip](../gossip/README.md)
- [distbuild](../distbuild/README.md)
- [wasm](../wasm/README.md)
