# About
A Web application to broadcast message on WhatsApp. This app uses the WhatsMeow package to connect with WhatsApp

# Interface
- Login
![login](https://github.com/Ganes556/wa-broadcast/blob/master/doc/login.png?raw=true)

# Run Development
- Change `.env` variable `ENVIRONTMENT` to "DEVELOPMENT"
- Go to `bundle` directory and run command `npm run dev`
- Go to root directory and run command `air .`

# Run production
- Change `.env` variable `ENVIRONTMENT` to "DEVELOPMENT"
- Go to `bundle` directory and run command `npm run build`
- Go to root directory and run command `go build -o <filename>`
- Run command `./<filename>.exe`

# Technology
- Go
- Fiber
- htmx.org
- Tailwind
- Daisyui
- Alpine.js