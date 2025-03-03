import sys
import argparse

from modules import converter

CMD_CONVERT = 1

def error(msg):
  print(f'error: {msg}')
  sys.exit(1)

def cmd_convert(subparser):
  parser = subparser.add_parser('convert',
                                aliases=['co'],
                                help='prepare audio files before transcribe')
  parser.add_argument('-s', '--src', required=False, type=str,
                      help='path to the folder with episodes audio files')
  parser.add_argument('-e', '--ep', required=False, type=str,
                      help='path to the show episode')
  parser.add_argument('-o', '--out', required=True, type=str,
                      help='path to the processed audio files')  
  parser.set_defaults(cmd=CMD_CONVERT)

def main():
  parser = argparse.ArgumentParser(prog='transcriber')
  subparsers = parser.add_subparsers(title='subcommands',
                                    help='additional help')
  # commands
  cmd_convert(subparsers)

  args = parser.parse_args(sys.argv[1:])
  if not hasattr(args, 'cmd'):
    error('subcommand is missing')

  if args.cmd == CMD_CONVERT:
    converter.main(args.src, args.ep, args.out)
  else:
    error('unknown subcommand')

main()
