module main

go 1.18

replace snipits/index => ./index

replace snipits/snipit => ./snipit

require snipits/index v0.0.0-00010101000000-000000000000

require snipits/snipit v0.0.0-00010101000000-000000000000 // indirect
