package blockchain

import (
	"work_queue"

	)

type miningWorker struct {
	// TODO. Should implement work_queue.Worker
	blk Block
	start uint64
	end uint64
}

func (workingWorker miningWorker) Run() interface{} {

	m := new(MiningResult)
	m.Found = false

	for i := workingWorker.start; i <= workingWorker.end; i++ {

		workingWorker.blk.Proof = i
		workingWorker.blk.Hash = workingWorker.blk.CalcHash()

		if workingWorker.blk.ValidHash() {

			m.Proof = i
			m.Found = true

			return *m

		}
	}
	return *m
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	// TODO
	//fmt.Println("work queue created")
	work_q := work_queue.Create(uint(workers), uint(chunks))

	mineResult := new(MiningResult)
	mineResult.Found = false

	chunkSize := (end - start) / chunks


	if chunkSize == 0 {
		chunkSize = end
	}

	for i := start; i < end; i += chunkSize {

		mineWorker := new(miningWorker)
		mineWorker.start = i
		if i+chunkSize-1 > end{
			mineWorker.end = end
		}else{
			mineWorker.end = i + chunkSize-1
		}
		mineWorker.blk = blk


		work_q.Enqueue(mineWorker)
	}

	// adapted from prof test cases
	running := true
	for running {

		 r := <- work_q.Results
			//fmt.Println("r.Found: ", r.(MiningResult).Found)
			if r.(MiningResult).Found == true {

				work_q.Shutdown()
				mineResult.Found = true
				mineResult.Proof = r.(MiningResult).Proof
				running = false
			}


	}

	return *mineResult
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << (8 * blk.Difficulty)) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}