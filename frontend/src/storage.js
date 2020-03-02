export function GetID() {
    return parseInt(sessionStorage.getItem('ID'));
}

export function SetID(id) {
    sessionStorage.setItem('ID', id);
}

function checkKey(n) {
    switch (n) {
        case 'Token':
        case 'Name':
        case 'Room':
        case 'Status':
            break;
        default:
            console.warn('Invalid storage key');
            break;
    }
}

export function GetStorage(n) {
    checkKey(n)
    return sessionStorage.getItem(n);
}

export function SetStorage(n, v) {
    checkKey(n)
    sessionStorage.setItem(n, v);
}

export function Clear() {
    sessionStorage.clear();
}