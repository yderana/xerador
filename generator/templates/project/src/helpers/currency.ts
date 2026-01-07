export function getFormatNumber(props: { numberValue: string, currencyCode?: any, comas?: number, enableComa?: boolean, visibleCurrencyCode?: boolean }) {
    const { numberValue, currencyCode, enableComa, visibleCurrencyCode } = props;
    let currency, localString, addon;
    const number = parseFloat(numberValue) ? numberValue.toString() : "0";
    const comaData = number.split('.')[1];
    switch (currencyCode) {
        case "IDR":
            currency = "Rp. ";
            localString = "id-ID";
            addon = comaData ? comaData : ",00";
            break;
        default:
            currency = "";
            localString = "en-En";
            addon = comaData ? comaData : ".00";
    }
    return (
        (visibleCurrencyCode ? "" : currency) + `${parseFloat(number).toLocaleString(localString)}` + (enableComa == false ? "" : addon)
    );
}