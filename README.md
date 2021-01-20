# Try To Stop It (BSOD provoking malware)

> <b>Attention!</b>
Use of the code samples and proof-of-concepts shown here is permitted solely at your own risk for academic and non-malicious purposes. It is the end user's responsibility to comply with all applicable local, state, and federal laws. The developer assumes no liability and is not responsible for any misuse or damage caused by this tool and the software in general. 

## What it is and how it works
The malware exploits a known bug in Windows 10, which provokes a bluescreen of death. In addition, the malware will create an entry in the Autostart menu and launch itself at every system startup. This is followed by an endless loop of BSODs.

The bug is described here: https://www.bleepingcomputer.com/news/security/windows-10-bug-crashes-your-pc-when-you-access-this-location/