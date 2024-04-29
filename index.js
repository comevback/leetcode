const converter = new showdown.Converter();
const showMD = (file) => {
    fetch(file)
        .then(response => response.text())
        .then(text => {
            const html = converter.makeHtml(text);
            document.querySelector('.markdown').innerHTML = html;
        })
        .catch(err => console.error('Error loading the Markdown file:', err));
};

showMD('./Memos/1672-Richest-Customer-Wealth.md');