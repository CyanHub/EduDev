const http = require('http');
const fs = require('fs');
const path = require('path');
const moment = require('moment');

// 确保日志文件存在
if (!fs.existsSync('log.txt')) {
    fs.writeFileSync('log.txt', '');
}

const server = http.createServer((req, res) => {
    // 记录访问日志
    const logEntry = `[${moment().format('YYYY-MM-DD HH:mm:ss')}] ${req.url}\n`;
    fs.appendFile('log.txt', logEntry, (err) => {
        if (err) {
            console.error('Failed to write log:', err);
        }
    });

    // 处理静态资源请求
    let filePath = path.join(__dirname, 'public', req.url === '/' ? 'index.html' : req.url);
    
    // 确定文件类型以设置正确的Content-Type
    const extname = path.extname(filePath);
    let contentType = 'text/html';
    switch (extname) {
        case '.js':
            contentType = 'text/javascript';
            break;
        case '.css':
            contentType = 'text/css';
            break;
        case '.json':
            contentType = 'application/json';
            break;
        case '.png':
            contentType = 'image/png';
            break;
        case '.jpg':
            contentType = 'image/jpg';
            break;
        case '.ico':
            contentType = 'image/x-icon';
            break;
    }

    // 读取文件并返回响应
    fs.readFile(filePath, (err, content) => {
        if (err) {
            if (err.code === 'ENOENT') {
                // 文件不存在，返回404页面
                fs.readFile(path.join(__dirname, 'public', '404.html'), (err, notFoundContent) => {
                    if (err) {
                        // 如果404页面也不存在，则返回纯文本404
                        res.writeHead(404, { 'Content-Type': 'text/plain' });
                        res.end('404 Not Found');
                    } else {
                        res.writeHead(404, { 'Content-Type': 'text/html' });
                        res.end(notFoundContent, 'utf8');
                    }
                });
            } else {
                // 其他错误（如权限问题），返回500错误
                res.writeHead(500);
                res.end(`Server Error: ${err.code}`);
            }
        } else {
            // 文件存在，正常返回
            res.writeHead(200, { 'Content-Type': contentType });
            res.end(content, 'utf8');
        }
    });
});

const PORT = 3000;
server.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
});    