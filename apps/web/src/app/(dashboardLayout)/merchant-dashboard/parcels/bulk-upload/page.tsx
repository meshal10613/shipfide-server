
import { UploadCloud, FileSpreadsheet, Download } from 'lucide-react';
import { Button } from '@/components/ui/Button';

export default function BulkUploadPage() {
  return (
    <div className="max-w-3xl mx-auto space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Bulk Parcel Import</h1>
        <p className="text-xs text-slate-500">Upload CSV or Excel file to import hundreds of orders at once</p>
      </div>

      <div className="bg-white p-8 rounded-3xl border border-slate-200 shadow-sm text-center space-y-6">
        <div className="border-2 border-dashed border-slate-200 rounded-3xl p-10 hover:border-[#DD0033]/40 transition-colors cursor-pointer space-y-3">
          <UploadCloud className="w-12 h-12 text-[#DD0033] mx-auto" />
          <h3 className="text-sm font-bold text-[#404042]">Drag & Drop your order file here</h3>
          <p className="text-xs text-slate-400">Supports .CSV, .XLSX (Max file size 10MB)</p>
          <Button variant="outline" size="sm" className="mt-2">Browse Computer</Button>
        </div>

        <div className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100 text-xs">
          <div className="flex items-center gap-3">
            <FileSpreadsheet className="w-5 h-5 text-emerald-600" />
            <span className="font-semibold text-[#404042]">Download CSV Template</span>
          </div>
          <Button variant="ghost" size="sm" leftIcon={<Download className="w-4 h-4" />}>
            Download
          </Button>
        </div>
      </div>
    </div>
  );
}

