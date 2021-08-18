function isPalindrome(str) {
    str = str.replace(/[^0-9a-z]/gi, '').toLowerCase();
    let checker = false
    for (let index = 0; index < str.length / 2; index++) {
        if (str[index] !== str[str.length - 1 - index]) {
            checker = true
        }
    }
    if (!checker) {
        console.log('It is a palindrome')
    } else {
        console.log('It is not  a palindrome')
    }
    console.log(str)
}

function uniqueUnion(...arguments) {
    for (let i in arguments) {
        for (let j in arguments[i]) {
            if (arguments[0].indexOf(arguments[i][j]) < 0) {
                arguments[0].push(arguments[i][j])
            }
        }
    }
    return (arguments[0])
}

function seekAndDestroy(arr, ...args) {
    arrr = arr.filter(e => args.indexOf(e) == -1);
    console.log(arrr)
}



toSpinalCase("NgoDuyDuuy=-DuyTrai")
function toSpinalCase(inputStr) {
    inputStr = inputStr.replace(/[^0-9a-z]/gi,"");
    console.log(inputStr)
    result = []
    for (let i = 0; i < inputStr.length; i++) {
        if (inputStr.charAt(i) === inputStr.charAt(i).toUpperCase()) {
            if (i==0) {
                result.push(inputStr[i])
            }else{
                result.push("-")
                result.push(inputStr[i])
            }
        }else{
            result.push(inputStr[i])
        }
    }
    console.log(result.join(""));
}

function drop(inputArr, callback) {
    outputArr=[]
    for (let i = 0; i < inputArr.length; i++) {
        if (callback(inputArr[i])==true) {
            outputArr.push(...inputArr.slice(i,))
            break
        }
    }
    console.log(outputArr);
}
drop([4,24,89,78,5,24], function (a) { return a >24; })
