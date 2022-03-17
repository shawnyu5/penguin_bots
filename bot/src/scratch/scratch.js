let { spawn } = require("child_process");
let exec = require("child_process").exec;

let output = spawn("python3", [
   "/home/shawn/python/penguin_bots/bot/src/scratch/scratch.py",
]);

output.stdout.on("data", (data) => {
   // console.log("(anon) data: %s", data); // __AUTO_GENERATED_PRINT_VAR__
});

exec(
   "python3 /home/shawn/python/penguin_bots/bot/src/scratch/scratch.py",
   (err, stdout, stderr) => {
      console.log("(anon) stdout: %s", stdout); // __AUTO_GENERATED_PRINT_VAR__
   }
);
