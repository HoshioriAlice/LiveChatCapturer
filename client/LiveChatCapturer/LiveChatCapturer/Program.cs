using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Threading;
using System.IO;
using System.Collections;
using System.Net;
using System.Net.Sockets;
using System.Net.Http;
using System.Text;
using System.Text.RegularExpressions;
namespace LiveChatCapturer
{
    static class Program
    {
        /// <summary>
        /// 应用程序的主入口点。
        /// </summary>
      
        [STAThread]
        static void Main()
        {
            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);
            var displayForm = new DisplayForm();
            var client = new Client("config.txt", displayForm.CommentBox);
            var th = new Thread(new ThreadStart(() => client.Start()));
            th.Start();
            Application.Run(displayForm.Form);
        }
    }
}
