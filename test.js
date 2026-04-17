const express = require("express")
const cors = require("cors")

const app = express()
app.use(cors())

let dataStore = {}

app.get("/set", (req, res) => {
  const key = req.query.key
  const value = req.query.value
  if (!key) return res.send("error")
  dataStore[key] = value
  res.send("saved")
})

app.get("/get", (req, res) => {
  let output = ""
  for (let k in dataStore) {
    output += k + ":" + dataStore[k] + "\n"
  }
  res.send(output || "empty")
})

app.get("/delete", (req, res) => {
  const key = req.query.key
  delete dataStore[key]
  res.send("deleted")
})

app.listen(3000)
