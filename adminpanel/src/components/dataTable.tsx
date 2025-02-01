import { useState } from "react";
import { DataTable } from "primereact/datatable";
import { Column } from "primereact/column";

interface Gorm {
  ID: number;
  CreatedAt?: string;
  UpdatedAt?: string;
  DeletedAt?: string | null;
}
interface Worker extends Gorm {
  name: string;
  barcode: string;
  position?: string;
}
interface Crop extends Gorm {
  name: string;
  season: string;
  type?: string;
}
interface Harvest extends Gorm {
  weight: number;
  bin: string;
  crop?: Crop;
  crop_id: number;
}
interface Product extends Gorm {
  name: string;
  type: string;
  unit: string;
  price: number;
  quantity: number;
  weight: number;
  length: number;
  width: number;
  height: number;
}
interface Process extends Gorm {
  unit: string;
  quantity: number;
  weight: number;
  student_use: number;
  value_added: number;
  harvest?: Harvest;
  harvest_id: number;
  product?: Product;
  product_id: number;
}
interface Hours extends Gorm {
  start: string;
  duration: number;
  department?: string;
  task?: string;
  worker?: Worker;
  worker_id: number;
  harvest?: Harvest;
  harvest_id?: number;
  process?: Process;
  process_id?: number;
}
interface ColField {
  field: string;
  header: string;
}
interface dataTableProps {
  rowData: Hours[];
  colDefs: ColField[];
}

export default function RowTable(props: dataTableProps) {
  const [rowData, setRowData] = useState(props.rowData);
  const [colDefs, setColDefs] = useState(props.colDefs);

  return (
    <div>
      <DataTable value={rowData} tableStyle={{ minWidth: "10rem" }}>
        {colDefs.map((col, i) => (
          <Column key={col.field} field={col.field} header={col.header} />
        ))}
      </DataTable>
    </div>
  );
}
