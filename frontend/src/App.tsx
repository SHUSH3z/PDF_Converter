import { useState } from "react";
import { OpenFolderDialog, SelectOutputFolder, ProcessExcelFiles } from "../wailsjs/go/main/App";

function App() {
  const [inputFolder, setInputFolder] = useState<string>("");
  const [outputFolder, setOutputFolder] = useState<string>("");
  const [status, setStatus] = useState<string>("");

  async function handleSelectInput() {
    const folder = await OpenFolderDialog();
    if (folder) setInputFolder(folder);
  }

  async function handleSelectOutput() {
    const folder = await SelectOutputFolder();
    if (folder) setOutputFolder(folder);
  }

  async function handleProcess() {
    if (!inputFolder || !outputFolder) {
      setStatus("Por favor, selecione as duas pastas.");
      return;
    }

    setStatus("Convertendo arquivos...");
    try {
      await ProcessExcelFiles(inputFolder, outputFolder);
      setStatus("Conversão concluída com sucesso!");
    } catch (err) {
      console.error(err);
      setStatus("Erro ao converter arquivos.");
    }
  }

  return (
    <div style={{
      minHeight: "100vh",
      minWidth: "100vw",
      backgroundColor: "#f0f2f5",
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
      justifyContent: "center",
      fontFamily: "'Poppins', sans-serif",
      padding: "2rem"
    }}>
      <h1 style={{
        fontSize: "2.5rem",
        marginBottom: "2rem",
        color: "#333",
        textAlign: "center"
      }}>
        🚀 SUPER CONVERSOR <br /> EXCEL ➔ PDF
      </h1>

      <div style={{ width: "100%", maxWidth: "500px", marginBottom: "1.5rem" }}>
        <button
          onClick={handleSelectInput}
          style={buttonStyle}
        >
          Selecionar Pasta de Entrada
        </button>
        {inputFolder && (
          <div style={folderPathStyle}>{inputFolder}</div>
        )}
      </div>

      <div style={{ width: "100%", maxWidth: "500px", marginBottom: "1.5rem" }}>
        <button
          onClick={handleSelectOutput}
          style={buttonStyle}
        >
          Selecionar Pasta de Saída
        </button>
        {outputFolder && (
          <div style={folderPathStyle}>{outputFolder}</div>
        )}
      </div>

      <div style={{ width: "100%", maxWidth: "500px", marginBottom: "2rem" }}>
        <button
          onClick={handleProcess}
          style={{ ...buttonStyle, backgroundColor: "#4caf50" }}
        >
          Iniciar Conversão
        </button>
      </div>

      <div style={{
        backgroundColor: "#fff",
        padding: "1rem 2rem",
        borderRadius: "8px",
        boxShadow: "0 4px 8px rgba(0,0,0,0.1)",
        color: "#333",
        fontSize: "1rem",
        textAlign: "center",
        maxWidth: "500px",
        width: "100%"
      }}>
        Status: <strong>{status || "Aguardando ação..."}</strong>
      </div>
    </div>
  );
}

const buttonStyle: React.CSSProperties = {
  width: "100%",
  padding: "0.8rem",
  backgroundColor: "#007bff",
  color: "white",
  border: "none",
  borderRadius: "8px",
  fontSize: "1rem",
  cursor: "pointer",
  transition: "background-color 0.3s ease",
};

const folderPathStyle: React.CSSProperties = {
  marginTop: "0.5rem",
  fontSize: "0.9rem",
  color: "#555",
  wordBreak: "break-all",
};

export default App;
