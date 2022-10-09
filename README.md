<h1 align="center">
Traffic Shadowing
</h1>

Using Goreplay to shadow http service traffic. (with Docker container example)

sudo gor --input-raw :8080 --output-http http://localhost:8081
