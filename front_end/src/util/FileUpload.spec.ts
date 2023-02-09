import { TestBed } from "@angular/core/testing";
import { pathArrayToDirStructure } from "./FileUpload";

// a test file for {fileName}/ts has the title {fileName}.spec.ts

describe('FileUpload basic', () => {
    let fileList1: string[] = [
        "f1.txt",
        "f2.txt",
    ]

    it('should work for empty input', () => {
        expect(
            pathArrayToDirStructure([])
        ).toEqual({})
    })
    
    it('should deal with flat files', () => {
        expect(
            pathArrayToDirStructure(fileList1)
        ).toEqual({
            "f1.txt": "file",
            "f2.txt": "file",
        })
    })

    let fileList2: string[] = [
        "a/f1.txt",
        "a/f2.txt",
    ]

    it('should deal with one directory', () => {
        expect(
            pathArrayToDirStructure(fileList2)
        ).toEqual({
            "a": {
                "f1.txt": "file",
                "f2.txt": "file",
        }})
    })
})
