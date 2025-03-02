import sys
import argparse

from modules import prepare

def get_args():
  parser = argparse.ArgumentParser(prog='SoundPreprocessor',
                                  description='Prepare audio files for transcribe')

  # define cli arguments
  parser.add_argument('-s', '--src', required=False, type=str,
                      help="path to the folder with episodes audio files")
  
  parser.add_argument('-e', '--ep', required=False, type=str,
                      help="path to the show episode")
  
  parser.add_argument('-o', '--out', required=True, type=str,
                      help="path to the processed audio files")

  args = parser.parse_args(sys.argv[1:])

  if (args.src == None and args.ep == None) or (args.src != None and args.ep != None):
    print("error: one of source parameters should be present")
    sys.exit(1)
  
  return args

def main():
  args = get_args()

  # prepare list of files for preprocessing
  files = [args.ep] if args.ep != None else prepare.get_episodes_path(args.src)
  print(f'start process `{len(files)}` file(s)')

  # start preprocessing
  prepare.preprocess(files, args.out)

main()

