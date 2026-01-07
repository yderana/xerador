export function generateIngridientSize(size: { length: number, width: number, height: number }, type: string) {
    let ingridientSize: any = { finalLength: null, finalWidth: null };
    switch (type.trim()) {
        case "Roll End Tuck Front":
            ingridientSize.finalLength = (2 * size.width) + (3 * size.height) + 10;
            ingridientSize.finalWidth = size.length + (4 * size.height) + 40;
            break;
        case "Roll End Tuck Top":
            ingridientSize.finalLength = (2 * size.width) + (2 * size.height) + 40;
            ingridientSize.finalWidth = size.length + (4 * size.height) + 40;
            break;
        case "Economic Tuck Top":
            ingridientSize.finalLength = (2 * size.width) + (3 * size.height) + 35;
            ingridientSize.finalWidth = size.length + (2 * size.height) + 10;
            break;
        case "Autobuttom Tuck Top":
            ingridientSize.finalLength = (2 * size.length) + (2 * size.width) + 35;
            ingridientSize.finalWidth = size.height + size.width + (1 / 2 * size.width) + 60;
            break;
        case "Reguler Slotted Container":
            ingridientSize.finalLength = 2 * (size.length + size.width) + 35;
            ingridientSize.finalWidth = size.width + size.height + 6;
            break;
        case "Full Overlap":
            ingridientSize.finalLength = 2 * (size.length + size.width) + 35;
            ingridientSize.finalWidth = (2 * size.width) + size.height;
            break;
        case "Goodie Box":
            ingridientSize.finalLength = 2 * (size.length + size.width) + 35;
            ingridientSize.finalWidth = size.length + size.width + 85;
            break;
        default:
            break;
    }
    return ingridientSize;
}