const express = require('express')
const app = express()

const PORT = 5000

app.get('/', (req, res) => {
  Promise.all([
  	Promise.resolve(fib(10)),
  	Promise.resolve(fib(20))
  ]).then(([fib10, fib20]) => {
    res.status(200).json({ fib10, fib20 })
  })
})
app.listen(PORT, () => {
  console.log(`listening to port *:${PORT}. press ctrl + c to cancel`)
})

function fib (n) {
  if (n <= 1) {
    return n
  }
  return fib(n - 1) + fib(n - 2)
}
