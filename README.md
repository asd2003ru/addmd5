# Вычисление md5 для ZMOD.

В архиве есть папка dist, в ней утилита вычисления и патча файла для использования с zmod

1. Выбираем под свою операционную систему:  

   - `addmd5_darwin_amd64`  (MacOS Intel)  
   - `addmd5_darwin_arm64` (MacOS Silicon)  
   - `addmd5_linux_amd64` (Linux)  
   - `addmd5_windows_amd64.exe` (Windows)  

2. Копируем выбранный файл куда нибудь.  
3. В Orca Slicer (Профиль процесса -> Прочее -> Скрипты постобработки) прописываем полный путь к файлу  

Например: `C:\addmd5_windows_amd64.exe`
