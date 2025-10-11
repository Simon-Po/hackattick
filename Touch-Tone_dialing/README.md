# Touchtone dialing

in this challenge what we pretty much want todo is decode a wav file.
While there are good decoders out there i have choosen todo it manually and get a bit into the bits

a good reference i found is 
https://www.videoproc.com/images/vp-seo/canonical-wav-file-structure.jpg

and the website itself https://www.videoproc.com/resource/wav-file.htm

i think for the beginning it makes sense to build a nice infrastructure in go to read specific amounts of bytes and convert them to the right types, especially because for now we dont know in what format the wav file is encoded, so we will just read the bytes and can then see that from the header.

Fuck this was a really hard one. In the end i solved it by translating this python code 

https://wirelesspi.com/goertzel-algorithm-evaluating-dft-without-dft/

(and a little help because i had no idea how to use math in go, by chatgpt)

then it was all about playing around with values and figuring out how to actually get the correct results, from time to time it felt a bit like shooting into the dar. 

while testing this, i also wrote a quick python script which can be used to generate dtmf signals and test against. 


