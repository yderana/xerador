const excel = require("node-excel-export");
// Create the excel report.
// This function will return Buffer
export function generateExcel(props: { name: string, heading: any, merges: any, specification: any, data: any[] }) {
    const { name, heading, merges, specification, data } = props;
    return excel.buildExport(
        [ // <- Notice that this is an array. Pass multiple sheets to create multi sheet report
            {
                name, // <- Specify sheet name (optional)
                heading, // <- Raw heading array (optional)
                merges, // <- Merge cell ranges
                specification, // <- Report specification
                data // <-- Report data
            }
        ]
    );
}