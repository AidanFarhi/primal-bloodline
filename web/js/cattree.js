let family = new FamilyTree(document.getElementById("tree"), {
    mode: 'dark',
    mouseScrool: FamilyTree.none,
    nodeBinding: {
        field_0: "name"
    }
})

// family.onUpdateNode((args) => {
// //post the data from args to your server
// })

// family.editUI.on('element-btn-click', function (sender, args) {
//     FamilyTree.fileUploadDialog(function (file) {
//         let formData = new FormData()
//         formData.append('file', file)
//         alert('upload the file')
//         console.log(args)
//     })
// })

family.load([
    { id: 1, pids: [2], name: "Nabi", gender: "female" },
    { id: 2, pids: [1], name: "Don Pedro", gender: "male" },
    { id: 3, mid: 1, fid: 2, name: "Killer", gender: "male" }
])