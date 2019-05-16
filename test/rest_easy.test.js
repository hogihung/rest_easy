const util = require('util');
const exec = util.promisify(require('child_process').exec);
const { mtrim } = require('js-trim-multiline-string');

const logFlag = 'N' // Y / N

test('use echo true command to verify exec child_process working correctly', async (done) => {
  const { stdout, stderr } = await exec('echo "true"');
  if (logFlag === 'Y') { console.log('stdout: ',stdout,'\n','stderr: ',stderr); }
  expect(stdout).toBe(mtrim(`true
    `))
  done();
});

test('command1 rest_easy', async (done) => {
  const { stdout, stderr } = await exec('rest_easy');
  if (logFlag === 'Y') { console.log('stdout: ',stdout,'\n','stderr: ',stderr); }
  expect(stdout).toBe( mtrim
    `REST Easy is a command line utility, which takes a JSON formatted configuration
     file and performs REST GET requests against the defined target endpoints. 

     Using this app, with JSON formatted config file, one can run n number of GET requests to the
     defined target endpoints and display the response to the terminal (default) and/or write the
     responses to a file.

     Usage:
       rest_easy [command]

     Available Commands:
       adhoc       Use the 'adhoc' sub-command to run GET requests against single endpoint
       help        Help about any command
       list        Use the 'list' sub-command to list the defined endpoints to be tested
       test        Use the 'test' sub-command to run GET requests to defined endpoints

     Flags:
       -h, --help         help for rest_easy
           --log string   log file (default is ./rest_easy.log)

     Use "rest_easy [command] --help" for more information about a command.
    `
  )
  done();
});

test('command2 rest_easy adhoc', async (done) => {
  try {
  const { stdout, stderr, err } = await exec('rest_easy adhoc');
  if (logFlag === 'Y') { console.log('stdout: ',stdout,'\n','stderr: ',stderr); }
  expect(stdout).toEqual(
    expect.stringContaining(
    "File does not exist:"
    )
  )
  }
  catch (e){
    console.log({e})
  }
  done();
});

test('command3', async (done) => {
  const { stdout, stderr } = await exec('echo "true"');
  if (logFlag === 'Y') { console.log('stdout: ',stdout,'\n','stderr: ',stderr); }
  expect(stdout).toBe( mtrim
    `true
    `
  )
  done();
});

test('command4', async (done) => {
  const { stdout, stderr } = await exec('echo "true"');
  if (logFlag === 'Y') { console.log('stdout: ',stdout,'\n','stderr: ',stderr); }
  expect(stdout).toBe( mtrim
    `true
    `
  )
  done();
});

test('command5', async (done) => {
  const { stdout, stderr } = await exec('echo "true"');
  if (logFlag === 'Y') { console.log('stdout: ',stdout,'\n','stderr: ',stderr); }
  expect(stdout).toBe( mtrim
    `true
    `
  )
  done();
});

test('command6', async (done) => {
  const { stdout, stderr } = await exec('echo "true"');
  if (logFlag === 'Y') { console.log('stdout: ',stdout,'\n','stderr: ',stderr); }
  expect(stdout).toBe( mtrim `true
    `
  )
  done();
});

test('verify folder structure prior to @availity/workflow scaffold', async (done) => {
  const { stdout, stderr } = await exec('ls');
  if (logFlag === 'Y') { console.log('stdout: ',stdout,'\n','stderr: ',stderr); }
  expect(stdout).toBe( mtrim `LICENSE
     README.md
     cmd
     main.go
     node_modules
     package-lock.json
     package.json
     rest_easy
     rest_easy.log
     targets.json
     test
    `
  )
  done();
});
