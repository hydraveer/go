import express from "express"

const app = express()
const port = 3000
app.use(express.json())
app.use(express.urlencoded({extended:true}))


app.get("/",(req,res)=>{
    res.send("Hi there")
})
app.get("/get",(req,res)=>{
    res.send("Hi there form inside")
})
app.post("/post",(req,res)=>{
    res.send(req.body)
})

app.listen(port,(req,res)=>{
    console.log(`Server is running on port ${port}`);
})