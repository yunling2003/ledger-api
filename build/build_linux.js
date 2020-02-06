var shell = require('shelljs');
shell.config.verbose = true
process.env.port  = 8080
process.env['GOOS'] = 'linux'
shell.rm('-rf', './dist')
shell.mkdir('dist')
shell.exec('go build -o ./dist/ledger-api ./src')
shell.cp('-rf', './dist/ledger-api', '../ledger-app/api/dist/')
