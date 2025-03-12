import sys
import os
import json

import soundfile as sf
from transformers import WhisperProcessor,AutoModelForSpeechSeq2Seq

MODEL = 'mitchelldehaven/whisper-medium-uk'

# initiate whisper model 
processor = WhisperProcessor.from_pretrained(MODEL)
model = AutoModelForSpeechSeq2Seq.from_pretrained(MODEL)

def get_files(path):
  files = []
  try:
    for file in os.listdir(path):
      # check file format to skip non audio
      if file.endswith('.wav'):
        files.append(os.path.join(path, file))
  except Exception as e:
    print(f'error: failed to list folder {path}:\n{e}')
    sys.exit(1)

  return files

def get_transcribe_filename(path, out):
  return os.path.join(out, f'{os.path.splitext(os.path.basename(path))[0]}.json')

def transcribe(file, out):
  audio_input, sample_rate = sf.read(file)

  chunk_size = int(25 * sample_rate) # 25 sec
  overlap_size = int(1 * sample_rate) # 1 sec

  transcribes = {}
  offset = 0

  # todo : improve transcription quality by merging chunks
  # - investigate approach with transformers pipeline function
  # - investigate custom model configuration to provide timestamps

  print(f'process file `{file}`')
  while offset <= len(audio_input):
    chunk_start = offset if offset == 0 else offset - overlap_size
    chunk_end = offset + chunk_size + overlap_size

    print(f'\ttranscribe chunk {int(chunk_start / sample_rate)}:{int(chunk_end / sample_rate)}s')
    chunk = audio_input[chunk_start:chunk_end]

    inputs = processor(chunk, sampling_rate=sample_rate, return_tensors='pt', return_attention_mask=True)
    predicted_ids = model.generate(inputs.input_features, attention_mask=inputs.attention_mask)

    result = processor.batch_decode(predicted_ids, skip_special_tokens=True)
    transcribes[f'{chunk_start}:{chunk_end}'] = result[0]

    offset += offset + chunk_size

  filepath = get_transcribe_filename(file, out)
  with open(filepath, 'w', encoding='utf-8') as file:
    json.dump(transcribes, file, ensure_ascii=False, indent=2)

  print(f'saved transcription into file `{filepath}`')
  

def main(src, ep, out):
  if (src == None and ep == None) or (src != None and ep != None):
      print("error: one of source parameters should be present")
      sys.exit(1)
  
  # prepare list of files for preprocessing
  files = [ep] if ep != None else get_files(src)
  print(f'start process `{len(files)}` file(s)')

  # start preprocessing
  for file in files:
    transcribe(file, out)
