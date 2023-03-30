# imago

🗺 Several image scaling algorithms implemented in Go.

## TODO

- [ ] Lanczos scaling
- [ ] Cubic scaling
- [ ] Rotating images
- [x] Convolutions / Filters

## Usage

After `go get -u github.com/fr3fou/imago/...`

```sh
$ lerp image.png 2
# or
$ nni image.png 2
```

~~hotel? trivago.~~

## References

- <http://courses.cs.vt.edu/~masc1044/L17-Rotation/ScalingNN.html>
- <https://www.wikiwand.com/en/Nearest-neighbor_interpolation>
- <https://www.wikiwand.com/en/Image_scaling>
- <https://www.wikiwand.com/en/Interpolation>
- <http://tech-algorithm.com/articles/bilinear-image-scaling/>
- <http://tech-algorithm.com/articles/linear-interpolation/>
- <https://www.wikiwand.com/en/Linear_interpolation>
- <https://www.cambridgeincolour.com/tutorials/image-interpolation.htm>
- <https://clouard.users.greyc.fr/Pantheon/experiments/rescaling/index-en.html>
- <https://www.wikiwand.com/en/Kernel_(image_processing)>
- <https://adeshpande3.github.io/A-Beginner%27s-Guide-To-Understanding-Convolutional-Neural-Networks-Part-2/>
- <https://www.codingame.com/playgrounds/2524/basic-image-manipulation/filtering>
- <https://www.researchgate.net/publication/276034240_Linear_Methods_for_Image_Interpolation>
- <http://web.pdx.edu/~jduh/courses/Archive/geog481w07/Students/Ludwig_ImageConvolution.pdf>
- <https://usman.it/image-manipulation-in-golang/>
- <https://ai.stanford.edu/~syyeung/cvweb/tutorial1.html>
- <https://setosa.io/ev/image-kernels/>

![](https://upload.wikimedia.org/wikipedia/commons/thumb/a/aa/Linear_interpolation_visualisation.svg/640px-Linear_interpolation_visualisation.svg.png)
![](https://upload.wikimedia.org/wikipedia/commons/1/19/2D_Convolution_Animation.gif)
