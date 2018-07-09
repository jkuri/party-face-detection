## party-face-detection

Experiment of face detection algorithm MTCNN ported to Golang on party people highly on drugs.

<p align="center">
  <img src="https://user-images.githubusercontent.com/1796022/42425963-c8d4232c-8323-11e8-8a13-1fca1fbd20a1.gif">
</p>

#### Running on MacOS

First, make sure you have all dependencies installed with running

```sh
brew install opencv go libtensorflow
```

Then, get the code

```sh
go -u -v github.com/jkuri/party-face-detection
```

...then, `cd $GOPATH/src/github.com/jkuri/party-face-detection` and build the app

```sh
make
```

At last, run `./build/party-face-detection` to process the video.

#### Command line parameters

```
Usage: party-face-detection [<flags>]

Flags:
      --help  Show context-sensitive help (also try --help-long and --help-man).
  -f, --file="./data/videos/ag.mp4"
              Input video file.
  -m, --model="./data/models/mtcnn.pb"
              MTCNN TF model file
```

#### Notes

Sample video is downloaded from YouTube [here](https://www.youtube.com/watch?v=4qMtqK8eTdM)
and is from parties in Ambasada Gavioli, Slovenia that lasted in 2009.

<p align="center">
  <img src="https://user-images.githubusercontent.com/1796022/42425971-eb367e24-8323-11e8-9fe2-f061c92d4d4f.png">
</p>

#### License

MIT
