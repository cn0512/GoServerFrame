taskkill /f /im Gate.exe
taskkill /f /im Game.exe
taskkill /f /im client.exe

start ./Instance/Gate/Gate.exe
start ./Instance/Game/Game.exe
start ./example/client/client.exe
