// This program test nodejs concurrency with promises

function delay1000 () {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve(true)
    }, 1000)
  })
}

function delay2000 () {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve(true)
    }, 2000)
  })
}

console.time('delay test')
Promise.all([delay1000(), delay2000()]).then((data) => {
  console.log(data)
  console.timeEnd('delay test')
})
