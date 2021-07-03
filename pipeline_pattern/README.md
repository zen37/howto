https://ketansingh.me/posts/pipeline-pattern-in-go-part-1/

Pipeline Pattern in Go Part 1
Posted at — Jul 2, 2021
I am going to divide this post into two parts, In the first part I will try to explain basic building blocks of pipelines and in second post I will try to build a general purpose library around this design.

I was recently reading Concurrency In Go and came across something called as pipeline processing pattern. Idea is you can break a logical functionality into stages. Each stage does its own processing and passes the output to the next stage to get processed. You can modify stages independent of one another, rate limit the stages and so on and forth.