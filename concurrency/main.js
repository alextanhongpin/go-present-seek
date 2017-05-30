// This program test nodejs concurrency with promises

function delay100 () {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve(true)
    }, 1000)
  })
}

function delay200 () {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve(true)
    }, 2000)
  })
}

console.time('delay test')
Promise.all([delay100(), delay200()]).then((data) => {
  console.log(data)
  console.timeEnd('delay test')
})
