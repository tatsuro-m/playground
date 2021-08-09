let count = parseInt(await $`ls -1 | wc -l`)
console.log(`Files count: ${count}`)

await $`docker --version`
await $`aws --version`

const displayWord = (word) => {
  console.log(`受け取った word: ${word}`);
}

displayWord("zx!!!!!!!!");

const getRandomWord = () => {
  return Math.random().toString(32).substring(2);
}

await $`echo ${getRandomWord()}`

const string = await $`cat README.md`
console.log(`READMEの中身:\n ${string}`)

const getDockerVersion = () => {
  return $`docker --version`;
}

$`echo Docker のバージョンは、${await getDockerVersion()}です！`
