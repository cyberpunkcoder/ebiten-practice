package game

var ()

func init() {
	/*
		audioContext, err = audio.NewContext(sampleRate)
		rcsFile, err := ebitenutil.OpenFile("assets/rcs.wav")

		rcsWav, err := wav.Decode(audioContext, rcsFile)

		rcs, err = audio.NewPlayer(audioContext, rcsWav)

		rcs.SetVolume(1)

		if err != nil {
			log.Fatal(err)
		}*/
}

// UpdateSound to continue looping if needed
func UpdateSound() {
	/*
		if int(rcs.Current().Seconds()) == 2 {
			rcs.Seek(1000000000)
		}*/
}

func startThrusterSound() {
	/*
		rcs.Seek(0)
		rcs.Play()
	*/
}

func stopThrusterSound() {
	/*
		if rcs.Current().Seconds() < 3 {
			rcs.Seek(3000000000)
		}
	*/
}
