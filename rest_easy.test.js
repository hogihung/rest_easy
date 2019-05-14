const util = require('util');
const exec = util.promisify(require('child_process').exec);
const { mtrim } = require('js-trim-multiline-string');


test('use echo true command to verify exec child_process working correctly', async (done) => {
  const { stdout, stderr } = await exec('echo "true"');
  console.log('stdout: ',stdout,'\n','stderr: ',stderr);
  expect(stdout).toBe(
    `true
    `.mtrim()
  )
  done();
});

test('verify folder structure prior to @availity/workflow scaffold', async (done) => {
  const { stdout, stderr } = await exec('ls');
  console.log('stdout: ',stdout,'\n','stderr: ',stderr);
  expect(stdout).toBe( mtrim
    `LICENSE
     README.md
     cmd
     main.go
     node_modules
     package-lock.json
     package.json
     rest_easy.log
     rest_easy.test.js
     targets.json
    `
  )
  done();
});
