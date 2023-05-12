
# distrubuted-chessboard-generation

A distributed system for generating all potential moves of a chessboard starting with 3 pieces.  Moves are to connected in a graph in sqlite db

# Installation

DCG requires golang, gRPC, and Proto 3.

# Usage

The master node will create a starter chessboard for all worker nodes that attempt to contact it.

Inside cb_master/, the master node must first be built and then ran:

    go build
    ./cb_master

Likewise in cb_worker/, the worker node must first be built and then ran.

Running cb_master will print an address that can then be used for the worker:

    go build
    ./cb_worker -masterAddr="PUTADDRESSHERE"

The worker asks for a starter board and then computes all possible moves for the recieved chessboard.

# To do
- [X] Implement basic chess rules
- [X] Generate starter boards with master
- [X] Send boards to workers
- [ ] Have workers send completed board tasks back to master as .db
- [ ] Implement checkmate/stalemate logic
- [ ] Add more tests to _test.go
