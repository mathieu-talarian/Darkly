let http = require("http");

let url = "http://192.168.186.129/.hidden/";

const regex = /<a href="(.+)">/gm;

let GetUrl = url => {
  http.get(url, res => {
    res.setEncoding("utf8");
    let rawData = "";
    let body = [];
    res.on("data", chunk => {
      body.push(chunk);
    });
    res
      .on("end", () => {
        let b = body.concat().toString();
        let m;
        while ((m = regex.exec(b)) !== null) {
          if (m.index === regex.lastIndex) {
            regex.lastIndex++;
          }
          m.forEach((elem, b) => {
            if (b === 1) {
              if (elem === "README") {
                  setTimeout(() => {
                    GetReadme(url + elem, rd => {
                        console.log("rd =>", rd)
                    }, 2000);
                  })
              } else if (elem === "../") {} 
              else GetUrl(url+elem)
            }
          });
        }
      })
      .on("error", e => {
        console.error(`Got error: ${e.message}`);
      });
  });
};

GetUrl(url);
let GetReadme = (url, cb) => {
    let body = [];
    http.get(url, res => {
        res
        .on("data", chunk => {
            body.push(chunk)
        })
        .on("end", () => {
            cb(body.concat().toString())
        })
        .on("error", () => {
            console.error(`error`)
        })
    })
}
