var shell = require('shelljs');
shell.config.verbose = true
process.env.port  = 8080
process.env['GOOS'] = 'linux'
shell.rm('-rf', './dist')
shell.mkdir('dist')
shell.cp('-rf', './config', './dist/')
shell.exec('go build -o ./dist/ledger-api ./src')
shell.cp('-rf', './dist/ledger-api', '../ledger-app/go/app/')
shell.cp('-rf', './dist/config/', '../ledger-app/go/app/')