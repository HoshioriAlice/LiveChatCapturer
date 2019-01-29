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
using System.Text;
using System.Text.RegularExpressions;
namespace LiveChatCapturer
{
    static class Program
    {
        /// <summary>
        /// 应用程序的主入口点。
        /// </summary>
        static void UpdateComment()
        {
            string liveURL;
            string serverIP;
            string serverPort;
            var filterStrings = new ArrayList();
            using(var reader = new StreamReader("config.txt"))
            {
                liveURL = reader.ReadLine();
                serverIP = reader.ReadLine();
                serverPort = reader.ReadLine();
                string s;
                while((s = reader.ReadLine()) != null)
                {
                    filterStrings.Add(s);
                }
            }
            var filter = new ArrayList();
            foreach (string filterString in filterStrings)
            {
                filter.Add(new Regex(filterString));
            }
            var client = new TcpClient(serverIP, Convert.ToInt32(serverPort));
            var buf = Encoding.ASCII.GetBytes(liveURL);
            var stream = client.GetStream();
            stream.Write(buf, 0, buf.Length);
            buf = new Byte[1024 * 16];
            int off = 0;
            while(true)
            {
                var count = stream.Read(buf, off, buf.Length);
                var message = Encoding.UTF8.GetString(buf, off, count);
                var commentList = message.Split('\n');
                for (int i = 0; i < commentList.Length - 1; ++i)
                {
                    bool isBlock = false;
                    foreach (Regex reg in filter)
                    {
                        if (reg.Matches(commentList[i]).Count != 0)
                        {
                            isBlock = true;
                        }
                    }
                    if (!isBlock) { commentBox.AppendText(commentList[i] + "\r\n"); }
                }
                for (int i = 0; i < buf.Length; ++i)
                {
                    buf[i] = 0;
                }
                for (int i =0;i < commentList[commentList.Length - 1].Length; ++i)
                {
                    buf[i] = Convert.ToByte(commentList[commentList.Length - 1][i]);
                }
                off = commentList[commentList.Length - 1].Length;
            }
        }
        [STAThread]
        static void Main()
        {
            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);
            var form = new Form1();
            //form.AutoScrollMinSize = new System.Drawing.Size(form.ClientRectangle.Width,form.ClientRectangle.Height);
            Control.CheckForIllegalCrossThreadCalls = false;
            commentBox = new TextBox();
            commentBox.Multiline = true;
            commentBox.Width = form.Width;
            commentBox.Height = form.Height;
            form.Resize += (_, __) =>
            {
                commentBox.Width = form.Width ;
                commentBox.Height = form.Height ;
            };
            form.Controls.Add(commentBox);
            var th = new Thread(new ThreadStart(UpdateComment));
            th.Start();
            Application.Run(form);
        }
        static private TextBox commentBox;
    }
}
