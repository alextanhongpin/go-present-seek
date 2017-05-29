const http = require('http')
const PORT = 3000

const handler = (req, res) => {
  Promise.all([
    Promise.resolve(fib(10)),
    Promise.resolve(fib(20))
  ]).then(([fib10, fib20]) => {
    res.writeHead(200, {'Content-Type': 'application/json'})
    res.end(JSON.stringify({ fib10, fib20 }))
  })
}

const server = http.createServer(handler)

server.listen(PORT, () => {
  console.log(`listening to port *:${PORT}. press ctrl + c to cancel`)
})

function fib (n) {
  if (n <= 1) {
    return n
  }
  return fib(n - 1) + fib(n - 2)
}
