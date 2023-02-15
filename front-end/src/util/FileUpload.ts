/**
 * Problem: given an aray of strings of the form:
 * "a/b/c/d/e.txt"..., build a "directory structure"
 * 
 */

// almost like defining our own var types
type myFile = string;
type myDirectory = {
    // the directory is a map from a name to either a file or another directory
    [name:string]: myFile | myDirectory
} 

export function pathArrayToDirStructure(files: string[]) : myDirectory {
    let root: myDirectory = {}

    files.forEach(
        file => {
            let path = file.split("/"); // array of all elements between slashes
            let fNameArr = path.splice(path.length-1); // splice out last entry
            let fName = fNameArr[0] // save last entry as the file name

            // traverse directories starting with root
            let dir = root;
            path.forEach( pEl => {
                let newDir = root[pEl];
                if (typeof newDir == "object") {
                    dir = newDir
                } else if (!newDir) {
                    newDir = {};
                    dir[pEl] = newDir;
                    dir = newDir;
                }
            })
            // dir is now innermost directory 

            dir[fName] = "file";

            root[file]="file"
        }
    )

    return root
}