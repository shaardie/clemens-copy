# clemens-copy

Dieses Programm überprüft alle Dateien **in einem Ordner (keine Unterordner!)** und verschiebt bestimmte Dateien in einen Unterordner namens `delete` (oder einem benutzerdefinierten Namen).

## 🔍 Was macht das Programm?

Wenn es in einem Ordner z. B. eine Datei `bild.png` gibt, aber **keine** Datei `bild.jpg` oder `bild.jpeg`,  
dann wird `bild.png` in den Ordner `delete` verschoben.

Wenn jedoch `bild.jpg` **existiert**, dann bleibt `bild.png` unangetastet.

## 🧰 Vorbereitung

1. **Lade die Datei `clemens-copy.exe` herunter**.
2. **Lege sie in den Ordner**, den du überprüfen willst.
3. Öffne die **Eingabeaufforderung (cmd)**:
   - Windows-Taste drücken
   - `cmd` eingeben und Enter drücken
4. **Wechsle in den Ordner**, z. B.:

   ```cmd
   cd C:\Users\DeinName\Bilder


5. **Führe das Programm aus**:

   ```cmd
   clemens-copy.exe
   ```

   Oder mit Optionen:

   ```cmd
   clemens-copy.exe -path "C:\Pfad\zum\Ordner" -delete "zum_loeschen"
   ```

   Oder als **Testlauf ohne tatsächliche Änderungen**:

   ```cmd
   clemens-copy.exe -dry
   ```

## ⚙️ Kommandozeilenoptionen

| Option      | Beschreibung                                                                 |
|-------------|------------------------------------------------------------------------------|
| `-path`     | (Optional) Pfad zum Ordner, der überprüft werden soll. Standard: aktueller Ordner |
| `-delete`   | (Optional) Name des Ordners für verschobene Dateien. Standard: `delete`          |
| `-dry`      | (Optional) Simulationsmodus – zeigt nur an, welche Dateien verschoben würden.  |

## 💡 Beispiel

Du hast einen Ordner `C:\Bilder`, darin liegt `clemens-copy.exe`. Dann:

```cmd
clemens-copy.exe
```

Oder als Testlauf:

```cmd
clemens-copy.exe -dry
```

Nach dem echten Lauf findest du einen neuen Ordner `delete`, in dem alle Dateien liegen,  
**die keine gleichnamige `.jpg` oder `.jpeg`-Datei haben**.

## ❗ Hinweise

- Das Programm **löscht keine Dateien**, sondern **verschiebt** sie nur.
- Du kannst den `delete`-Ordner danach manuell prüfen oder löschen.
- Das Programm bearbeitet nur Dateien **im angegebenen Ordner**, **nicht in Unterordnern**.

---

Viel Erfolg!
